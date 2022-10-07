package main

import (
    "github.com/enzdor/ratings-api/routers"
    "github.com/enzdor/ratings-api/utils/database"

)

func main() {
    database.StartDatabase()
    routers.StartRouters()
}
