package main

import (
    "fmt"
    "users_crud/internal/user"
)

func main() {

    users := []internal.User{
        {
            ID:    12,
            Name:  "Israel L.",
            Email: "johndoe@mail.com",
        },
        {
            ID:    13,
            Name:  "Javier L.",
            Email: "joanede@mail.com",
        },
        {
            ID:    14,
            Name:  "Pako L.",
            Email: "juande@mail.com",
        },
    }

    result, err := internal.UpdateUser(
        users,
        13,
        "pakomail",
        "johndoe",
    )

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(*result)
}