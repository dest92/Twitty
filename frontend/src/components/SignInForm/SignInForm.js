import React from "react";
import "./SignInForm.scss";
import { Form, Button, Spinner } from "react-bootstrap";

export default function SignInForm() {
  const onSubmit = (e) => {
    e.preventDefault();
    console.log("Login");
  };

  return (
    <div className="signin-form">
      <h2>Get in</h2>
      <Form onSubmit={onSubmit}>
        <Form.Group>
          <Form.Control type="email" placeholder="Email Address" />
          <Form.Control type="password" placeholder="Password" />
        </Form.Group>
        <Button variant="primary" type="submit">
            Login
        </Button>
      </Form>
    </div>
  );
}
