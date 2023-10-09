import * as improvmx from "@pulumi/improvmx";

const random = new improvmx.Random("my-random", { length: 24 });

export const output = random.result;