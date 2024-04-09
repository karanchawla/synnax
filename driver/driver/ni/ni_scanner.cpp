//
// Created by Synnax on 3/24/2024.
//

#include "ni_scanner.h"
#include "nisyscfg.h"

ni::NiScanner::NiScanner() {}

ni::NiScanner::~NiScanner() {}

ni::NiScanner::json ni::NiScanner::getDevices() {
    json j;
    char productName[1024] = "";
    char serialNumber[1024] = "";
    char isSimulated[1024] = "";
    NISysCfgStatus status = NISysCfg_OK;
    NISysCfgEnumResourceHandle resourcesHandle = NULL;
    NISysCfgResourceHandle resource = NULL;
    NISysCfgFilterHandle filter = NULL;
    NISysCfgSessionHandle session = NULL;

    // initialized cfg session
    status = NISysCfgInitializeSession( //TODO: look into this
            "localhost",            // target (ip, mac or dns name)
            NULL,                   // username (NULL for local system)
            NULL,                   // password (NULL for local system)
            NISysCfgLocaleDefault,  // language
            NISysCfgBoolTrue,       //force pproperties to be queried everytime rather than cached
            10000,                  // timeout (ms)
            NULL,                   // expert handle
            &session                //session handle
    );
    // Attempt to find hardware
    NISysCfgFindHardware(session, NISysCfgFilterModeAll, filter, NULL, &resourcesHandle);
    j["devices"] = json::array();
    // Iterate through all hardware found and grab the relevant information
    while(NISysCfgNextResource(session, resourcesHandle, &resource)  == NISysCfg_OK) { // instead  do while (!= NISysCfgWarningNoMoreItems) ?
        json device;
        NISysCfgGetResourceProperty(resource, NISysCfgResourcePropertyProductName, productName);
        NISysCfgGetResourceProperty(resource, NISysCfgResourcePropertySerialNumber, serialNumber);
        NISysCfgGetResourceProperty(resource, NISysCfgResourcePropertyProductName, productName);
        NISysCfgGetResourceProperty(resource, NISysCfgResourcePropertyIsSimulated, isSimulated);
        device["productName"] = productName;
        device["serialNumber"] = serialNumber;
        device["isSimulated"] = isSimulated;
        j["devices"].push_back(device);
    }
    NISysCfgCloseHandle(resourcesHandle);
    return json;
}