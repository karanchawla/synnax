// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#pragma once

#include "client/cpp/telem/telem.h"
#include "x/go/telem/x/go/telem/telem.pb.h"
#include <string>
#include <vector>
#include <cstddef>

constexpr auto NEWLINE_TERMINATOR = static_cast<std::byte>('\n');

namespace synnax {
/// @brief Series is a strongly typed array of telemetry samples backed by an underlying binary buffer.
class Series {
public:
    Series(DataType data_type, size_t size): data_type(data_type),
                                             data(std::make_unique<std::byte[]>(
                                                 size * data_type.density())),
                                             size(size * data_type.density()) {
    }

    /// @brief constructs a series from the given vector of numeric data.
    template<typename NumericType>
    explicit Series(const std::vector<NumericType> &d,
                    DataType data_type_ = DATA_TYPE_UNKNOWN) : data_type(
        std::move(data_type_)) {
        static_assert(std::is_arithmetic_v<NumericType>,
                      "NumericType must be a numeric type");
        if (data_type == DATA_TYPE_UNKNOWN)
            data_type = DataType::from_type<NumericType>();
        size = d.size() * data_type.density();
        data = std::make_unique<std::byte[]>(size);
        memcpy(data.get(), d.data(), size);
    }

    explicit Series(TimeStamp v) : data_type(synnax::TIMESTAMP) {
        size = data_type.density();
        data = std::make_unique<std::byte[]>(size);
        memcpy(data.get(), &v.value, size);
    }

    /// @brief constructs a series of length 1 from the given number.
    template<typename NumericType>
    explicit Series(NumericType t, DataType data_type_ = DATA_TYPE_UNKNOWN): data_type(
        data_type_) {
        // single sample constructor
        static_assert(std::is_arithmetic_v<NumericType>,
                      "NumericType must be a numeric type");
        if (data_type == DATA_TYPE_UNKNOWN)
            data_type = DataType::from_type<NumericType>();
        size = data_type.density();
        data = std::make_unique<std::byte[]>(size);
        memcpy(data.get(), &t, size);
    }

    template<typename NumericType>
    void set(size_t index, NumericType value) {
        if (index >= size / data_type.density()) {
            throw std::runtime_error("index out of bounds");
        }
        memcpy(data.get() + index * data_type.density(), &value, data_type.density());
    }


    // set an array (not a vector)
    template<typename NumericType>
    void set_array(const NumericType *d, size_t index, size_t length) {
        if (index + length > size / data_type.density()) {
            throw std::runtime_error("index out of bounds");
        }
        memcpy(data.get() + index * data_type.density(), d,
               length * data_type.density());
    }


    /// @brief constructs the series from the given vector of strings. These can also
    /// be JSON encoded strings, in which case the data type should be set to JSON.
    /// @param d the vector of strings to be used as the data.
    /// @param data_type_ the type of data being used.
    explicit Series(
        const std::vector<std::string> &d,
        DataType data_type_ = STRING
    ) : data_type(std::move(data_type_)) {
        if (data_type != STRING && data_type != JSON)
            throw std::runtime_error("invalid data type b");
        size_t total_size = 0;
        for (const auto &s: d) total_size += s.size() + 1;
        data = std::make_unique<std::byte[]>(total_size);
        size_t offset = 0;
        for (const auto &s: d) {
            memcpy(data.get() + offset, s.data(), s.size());
            offset += s.size();
            data[offset] = static_cast<std::byte>('\n');
            offset++;
        }
        size = total_size;
    }

    /// @brief constructs the series from its protobuf representation.
    explicit Series(const telem::PBSeries &s) : data_type(s.data_type()) {
        size = s.data().size();
        data = std::make_unique<std::byte[]>(size);
        memcpy(data.get(), s.data().data(), size);
    }

    /// @brief encodes the series' fields into the given protobuf message.
    /// @param pb the protobuf message to encode the fields into.
    void to_proto(telem::PBSeries *pb) const {
        pb->set_data_type(data_type.name());
        pb->set_data(data.get(), size);
    }

    [[nodiscard]] std::vector<uint8_t> uint8() const {
        if (data_type != synnax::UINT8) {
            throw std::runtime_error("invalid data type");
        }
        std::vector<uint8_t> v(size);
        memcpy(v.data(), data.get(), size);
        return v;
    }

    [[nodiscard]] std::vector<float> float32() const {
        if (data_type != synnax::FLOAT32) {
            throw std::runtime_error("invalid data type");
        }
        std::vector<float> v(size / sizeof(float));
        memcpy(v.data(), data.get(), size);
        return v;
    }

    [[nodiscard]] std::vector<int64_t> int64() const {
        if (data_type != synnax::INT64) {
            throw std::runtime_error("invalid data type");
        }
        std::vector<int64_t> v(size / sizeof(int64_t));
        memcpy(v.data(), data.get(), size);
        return v;
    }

    [[nodiscard]] std::vector<uint64_t> uint64() const {
        if (data_type != synnax::UINT64 && data_type != synnax::TIMESTAMP) {
            throw std::runtime_error("invalid data type");
        }
        std::vector<uint64_t> v(size / sizeof(uint64_t));
        memcpy(v.data(), data.get(), size);
        return v;
    }

    [[nodiscard]] std::vector<std::string> string() const {
        if (data_type != synnax::STRING && data_type != synnax::JSON) {
            throw std::runtime_error("invalid data type");
        }
        std::vector<std::string> v;
        std::string s;
        for (size_t i = 0; i < size; i++) {
            if (data[i] == static_cast<std::byte>('\n')) {
                v.push_back(s);
                s.clear();
                // WARNING: This might be very slow due to copying.
            } else s += char(data[i]);
        }
        return v;
    }

    Series(const Series &s) : data_type(s.data_type) {
        data = std::make_unique<std::byte[]>(s.size);
        memcpy(data.get(), s.data.get(), s.size);
        size = s.size;
    }

    // implement the array access operator
    template<typename NumericType>
    NumericType operator[](int index) const {
        return at<NumericType>(index);
    }

    template<typename NumericType>
    NumericType at(int index) const {
        if (index >= length())
            throw std::runtime_error("index" + std::to_string(index) + " out of bounds for series");
        if (index < 0) index = static_cast<int>(length()) + index;
        NumericType value;
        memcpy(&value, data.get() + index * data_type.density(), data_type.density());
        return value;
    }

    size_t length() const {
        return size / data_type.density();
    }

    // implement the ostream operator
    friend std::ostream &operator<<(std::ostream &os, const Series &s) {
        os << "Series(" << s.data_type.name() << ", [";
        if (s.data_type == synnax::STRING || s.data_type == synnax::JSON) {
            auto strings = s.string();
            for (size_t i = 0; i < strings.size(); i++) {
                os << "\"" << strings[i] << "\"";
                if (i < strings.size() - 1) os << ", ";
            }
        } else if (s.data_type == synnax::FLOAT32) {
            auto floats = s.float32();
            for (size_t i = 0; i < floats.size(); i++) {
                os << floats[i];
                if (i < floats.size() - 1) os << ", ";
            }
        } else if (s.data_type == synnax::INT64) {
            auto ints = s.int64();
            for (size_t i = 0; i < ints.size(); i++) {
                os << ints[i];
                if (i < ints.size() - 1) os << ", ";
            }
        } else if (s.data_type == synnax::UINT64 || s.data_type == synnax::TIMESTAMP) {
            auto ints = s.uint64();
            for (size_t i = 0; i < ints.size(); i++) {
                os << ints[i];
                if (i < ints.size() - 1) os << ", ";
            }
        } else if (s.data_type == synnax::UINT8) {
            auto ints = s.uint8();
            for (size_t i = 0; i < ints.size(); i++) {
                os << ints[i];
                if (i < ints.size() - 1) os << ", ";
            }
        } else {
            os << "unknown data type";
        }
        os << "])";
        return os;
    }

    /// @brief Holds what type of data is being used.
    DataType data_type;

    /// @brief Holds the underlying data.
    std::unique_ptr<std::byte[]> data;

    /// @brief an optional property that defines the time range occupied by the Series' data. This property is
    /// guaranteed to be defined when reading data from a Synnax cluster, and is particularly useful for understanding
    /// the alignment of samples in relation to another series. When read from a cluster, the start of the time range
    /// represents the timestamp of the first sample in the array (inclusive), while the end of the time
    /// range is set to the nanosecond AFTER the last sample in the array (exclusive).
    synnax::TimeRange time_range = synnax::TimeRange();


    size_t size;
};
}
