import {email, maxLength, minLength, object, type Output, string} from "valibot";

export const schemaLoginRequest = object({
    email: string([email("Invalid email address.")]),
    password: string([minLength(8, "The password must be at least 8 characters long."), maxLength(256, "The password must be at most 256 characters long.")]),
})

export type LoginRequest = Output<typeof schemaLoginRequest>;

export const schemaLoginResponse = object({
    token: string(),
});

export type LoginResponse = Output<typeof schemaLoginResponse>;
