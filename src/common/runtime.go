package common

import "path/filepath"

var (
	PATH_ROOT string
	PATH_MODELES string
	PATH_ROUTERS string
	PATH_CONTROLLERS string
	PATH_VIEWS string
	PATH_SQL string
	PATH_JSON string

)

const (
	DIR_SRC = "src"
	DIR_SRC_MODELS = "models"
	DIR_SRC_ROUTERS = "routers"
	DIR_SRC_CONTROLLERS = "controllers"
	DIR_VIEW = "views"
	DIR_SQL = "sql"
	DIR_JSON = "json"
)

func init()  {
	PATH_ROOT, _ = filepath.Abs("./")
	PATH_MODELES = filepath.Join(PATH_ROOT, DIR_SRC, DIR_SRC_MODELS)
	PATH_ROUTERS = filepath.Join(PATH_ROOT, DIR_SRC, DIR_SRC_ROUTERS)
	PATH_CONTROLLERS = filepath.Join(PATH_ROOT, DIR_SRC, DIR_SRC_CONTROLLERS)
	PATH_VIEWS = filepath.Join(PATH_ROOT, DIR_VIEW)
	PATH_SQL = filepath.Join(PATH_ROOT, DIR_SQL)
	PATH_JSON = filepath.Join(PATH_ROOT, DIR_JSON)
}
