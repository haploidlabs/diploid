import { coerce, number } from "valibot";

export const schemaId = coerce(number(), Number);
