// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("improvmx");

/**
 * API key for the ImprovMX
 */
export declare const api_key: string | undefined;
Object.defineProperty(exports, "api_key", {
    get() {
        return __config.get("api_key");
    },
    enumerable: true,
});

export declare const version: string | undefined;
Object.defineProperty(exports, "version", {
    get() {
        return __config.get("version");
    },
    enumerable: true,
});

