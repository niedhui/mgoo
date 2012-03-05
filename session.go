package mgoo

import (
  "launchpad.net/mgo"
)

func WithMgo(url string, f func(*mgo.Session)) {
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  f(session)
}
