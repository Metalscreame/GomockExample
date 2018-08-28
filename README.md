# GomockExample
Its a repo that shows how to use Gomock framework

How To Generate a mock:

    mockgen -destination=dstPath/file_name.go -package=pkgNameToGenerate github.com/repo_name/project_name/src/dir_with_interface InterfaceName

To generate UserPersistance mock:

    mockgen -destination=src/mocks/mock_UserPersistence.go -package=mocks github.com/metalscreame/GomockExample/src/models UserPersistence


or run 

    go generate ./...
    
when there is a comment section in the interface file