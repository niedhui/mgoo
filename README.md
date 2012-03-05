some convenient function for [mgo](http://labix.org/mgo)

```
  //WithMgo will create a session ,execute the function,then close the session
  mgoo.WithMgo("localhost", func(session *mgo.Session) {
    coll := session.DB("hello_mgo").C("products")
    coll.RemoveAll(nil)
    coll.Insert(bson.M{"name": "cola", "price": 12})
  })
```