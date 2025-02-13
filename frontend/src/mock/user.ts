import { User } from "../types"

export const mockUser: User = {
    "name": "John Doe",
    "username": "johndoe123",
    "email": "john.doe@example.com",
    "role": "user",
    "memberSince": "2021-05-15",
    "addresses": [
        {
            "address": "123 Main St",
            "city": "Springfield",
            "zipCode": "12345"
        },
        {
            "address": "456 Elm St",
            "city": "Shelbyville",
            "zipCode": "67890"
        }
    ]
}
