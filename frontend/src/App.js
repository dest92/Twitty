import React from "react";
import { useState } from "react";
import { Button } from "react-bootstrap";

export default function App() {
  const [user, setUser] = useState();

  // return (
  //   <div>{user ? <h1>You´re logged</h1> : <h1>You´re not logged</h1>}</div>
  // );

  return (
    <div>
      <Button>Login</Button>
    </div>
  );
}
