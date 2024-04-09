#pragma once

#include <string>
#include "nlohmann/json.hpp"

using json = nlohmann::json;

/// @brief general utilities for parsing configurations.
namespace config {
/// @brief a utility class for improving the experience of parsing JSON-based
/// configurations.
class Parser {
public:
    std::shared_ptr<std::vector<json> > errors;

    /// @brief constructs a parser for accessing values on the given JSON configuration.
    explicit Parser(json config): errors(std::make_shared<std::vector<json> >()),
                                  config(std::move(config)) {
    }

    /// @brief constructs a parser for accessing values on the given stringified
    /// JSON configuration. If the string is not valid JSON, immediately binds an error
    /// to the parser.
    explicit Parser(const std::string &encoded): errors(std::make_shared<std::vector<json> >()) {
        try {
            config = json::parse(encoded);
        } catch (const json::parse_error &e) {
            noop = true;
            field_err("", e.what());
        }
    }

    /// @brief default constructor constructs a parser that will fail fast.
    Parser(): errors(nullptr), noop(true) {
    }


    /// @brief gets the field at the given path. If the field is not found,
    /// accumulates an error in the builder.
    template<typename T>
    T required(const std::string &path) {
        if (noop) return T();
        const auto iter = config.find(path);
        if (iter == config.end()) {
            field_err(path, "This field is required");
            return T();
        }
        return get<T>(path, iter);
    }

    /// @brief attempts to pull the value at the provided path. If that path is not found,
    /// returns the default. Note that this function will still accumulate an error if the
    /// path is found but the value is not of the expected type.
    /// @param path The JSON path to the value.
    /// @param default_value The default value to return if the path is not found.
    template<typename T>
    T optional(const std::string &path, T default_value) {
        if (noop) return default_value;
        const auto iter = config.find(path);
        if (iter == config.end()) return default_value;
        return get<T>(path, iter);
    }

    /// @brief gets the field at the given path and creates a new parser just for that
    /// field. The field must be an object or an array. If the field is not of the
    /// expected type, or if the field is not found, accumulates an error in the parser.
    /// @param path The JSON path to the field.
    Parser child(const std::string &path) const {
        if (noop) return {};
        const auto iter = config.find(path);
        if (iter == config.end()) {
            field_err(path, "This field is required");
            return {};
        }
        if (!iter->is_object() && !iter->is_array()) {
            field_err(path, "Expected an object or array");
            return {};
        }
        return {*iter, errors, path_prefix + path + "."};
    }

    /// @brief Iterates over an array at the given path, executing a function for each element.
    /// If the path does not point to an array, logs an error.
    /// @param path The JSON path to the array.
    /// @param func The function to execute for each element of the array. It should take a
    /// Parser as its argument.
    void iter(
        const std::string &path,
        const std::function<void(Parser &)> &func
    ) const {
        if (noop) return;
        const auto iter = config.find(path);
        if (iter == config.end())return field_err(path, "This field is required");
        if (!iter->is_array()) return field_err(path, "Expected an array");
        for (size_t i = 0; i < iter->size(); ++i) {
            const auto child_path = path_prefix + path + "." + std::to_string(i) + ".";
            Parser childParser((*iter)[i], errors, child_path);
            func(childParser);
        }
    }

    /// @brief binds a new error to the field at the given path.
    /// @param path The JSON path to the field.
    /// @param message The error message to bind.
    void field_err(const std::string &path, const std::string &message) const {
        if (noop) return;
        errors->push_back({
            {"path", path_prefix + path},
            {"message", message}
        });
    }

    /// @returns true if the parser has accumulated no errors, false otherwise.
    [[nodiscard]] bool ok() const {
        if (noop) return false;
        return errors->empty();
    }

    /// @returns the parser's errors as a JSON object of the form {"errors": [ACCUMULATED_ERRORS]}.
    [[nodiscard]] json error_json() const {
        json err;
        err["errors"] = *errors;
        return err;
    }

private:
    /// @brief the JSON configuration being parsed.
    json config;
    /// @brief used for tracking the path of a child parser.
    std::string path_prefix;
    /// @brief noop means the parser should fail fast.
    bool noop = false;

    Parser(
        json config,
        std::shared_ptr<std::vector<json> > errors,
        std::string path_prefix
    ): errors(std::move(errors)),
       config(std::move(config)),
       path_prefix(std::move(path_prefix)) {
    }

    template<typename T>
    T get(const std::string &path, const nlohmann::basic_json<>::iterator &iter) {
        try {
            return iter->get<T>();
        } catch (const nlohmann::json::type_error &e) {
            // slice the error message from index 32 to remove the library error prefix.
            field_err(path, e.what() + 32);
        }
        return T();
    }
};
}