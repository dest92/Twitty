import React, { useState } from "react";
import "./SignInForm.scss";
import { Form, Button, Spinner } from "react-bootstrap";
import { values, size } from "lodash";
import { toast } from "react-toastify";
import { isEmailValid } from "../../utils/validations";

export default function SignInForm() {
  const [formData, setFormData] = useState(initialForm(  ));
  const [signInLoading, setSignInLoading] = useState(false);
  const onSubmit = (e) => {
    e.preventDefault();

    if (!navigator.onLine) {
      toast.error("No internet connection");
      return;
    }

    let validCount = 0;

    values(formData).some((value) => {
      value && validCount++;
      return null;
    });

    if (size(formData) !== validCount) {
      toast.warning("Please fill out all fields");
      return;
    } else {
      if (!isEmailValid(formData.email)) {
        toast.warning("Please enter a valid email address");
      } else {
        setSignInLoading(true);
        toast.success("Success!");
      }
    }
  };

  const onChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  return (
    <div className="signin-form">
      <h2>Get in</h2>
      <Form onSubmit={onSubmit} onChange={onChange}>
        <Form.Group>
          <Form.Control
            type="email"
            placeholder="Email Address"
            name="email"
            defaultValue={formData.email}
          />
          <Form.Control
            type="password"
            placeholder="Password"
            name="password"
            defaultValue={formData.password}
          />
        </Form.Group>
        <Button variant="primary" type="submit">
          {!signInLoading ? "Login" : <Spinner animation="border" size="sm" />}
        </Button>
      </Form>
    </div>
  );
}

function initialForm() {
  return {
    email: "",
    password: "",
  };
}
