package mgoo

import (
  "launchpad.net/mgo"
  "launchpad.net/mgo/bson"
)

func WithMgo(url string, f func(*mgo.Session)) {
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  f(session)
}

func FindById(coll *mgo.Collection, id interface{}) (result bson.M) {
  result = make(bson.M)
  var queryId bson.ObjectId
  switch id := id.(type) {
  case string:
    queryId = bson.ObjectIdHex(id)
  case bson.ObjectId:
    queryId = id
  default:
  }
  coll.Find(bson.M{"_id": queryId}).One(result)
  return
}
