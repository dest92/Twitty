import React from "react";
import { useState } from "react";
import SignInSignUp from "./page/SignInSignUp/";
import { ToastContainer } from "react-toastify";

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
      <ToastContainer
        position="top-right"
        autoClose={5000}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        pauseOnHover
        
      />
    </div>
  );
}
