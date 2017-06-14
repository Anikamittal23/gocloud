package gocloud

import(
  "github.com/scorelab/gocloud-v2/aws"
  "github.com/scorelab/gocloud-v2/google"
   "fmt"
  "errors"
)

type Gocloud interface{
  Createnode(request interface{})(resp interface{},err error)
  Startnode(request interface{}) (resp interface{}, err error)
  Stopnode(request interface{}) (resp interface{}, err error)
  Deletenode(request interface{}) (resp interface{}, err error)
  Rebootnode(request interface{}) (resp interface{}, err error)
}

const(
  Amazonprovider = "aws"
  Googleprovider = "google"
  Secretkey = "dummy"
  Secretid = "dummy"
)


func CloudProvider(provider,Secretkey,secretid string)(Gocloud, error){

    switch provider{
    case Amazonprovider:
       return new(aws.AWS), nil

    case Googleprovider:
       return new(google.Google), nil

      default:
        return nil, errors.New(fmt.Sprintf("provider %s not recognized\n",provider))
    }
}
