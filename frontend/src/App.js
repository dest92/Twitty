import React from "react";
import { useState } from "react";
import SignInSignUp from "./page/SignInSignUp/";

export default function App() {
  const [user, setUser] = useState(true);

  return (
    <div>
      {user ? (
        <div>
          <SignInSignUp />
        </div>
      ) : (
        <h1>You´re not logged</h1>
      )}
    </div>
  );
}
