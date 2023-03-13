import React, { useState } from "react";
import "./SignUpForm.scss";
import { Row, Col, Form, Button, Spinner } from "react-bootstrap";
import { values, size } from "lodash";
import { toast } from "react-toastify";
import { isEmailValid } from "../../utils/validations";
import { signUpApi } from "../../api/auth";

export default function SignUpForm(props) {
  const [formData, setFormData] = useState(initialForm());
  const { setShowModal } = props;

  const [signUpLoading, setSignUpLoading] = useState(false);

  const onSubmit = (e) => {
    e.preventDefault(); //Prevents page refresh
    // console.log(formData);

    if (!navigator.onLine) {
      toast.error("No internet connection");
      return;
    }

    let validCount = 0;

    values(formData).some((value) => {
      value && validCount++;
      return null;
    });

    // console.log("Size: " + size(formData));

    if (validCount !== 5) {
      toast.warning("Please fill out all fields");
      return;
    } else {
      if (!isEmailValid(formData.email)) {
        toast.warning("Please enter a valid email address");
      } else if (formData.password !== formData.repeatPassword) {
        toast.warning("Passwords do not match");
      } else if (size(formData.password) < 8) {
        toast.warning("Password must be at least 8 characters");
      } else if (size(formData.name) < 2 || size(formData.lastName) < 2) {
        toast.warning("Name and last name must be at least 2 characters");
      } else {
        setSignUpLoading(true);
        signUpApi(formData)
          .then((response) => {
            if (response.code) {
              toast.warning(response.message);
            } else {
              toast.success("Success!");
              setShowModal(false);
              setFormData(initialForm());
            }
          })
          .catch(() => {
            toast.warning("Error from server, try again later!");
          })
          .finally(() => {
            setSignUpLoading(false);
          });
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
    <div className="sign-up-form">
      <h2>Create your account</h2>
      <Form onSubmit={onSubmit} onChange={onChange}>
        <Form.Group>
          <Row>
            <Col>
              <Form.Control
                type="text"
                name="name"
                placeholder="Name"
                defaultValue={formData.name}
              />
            </Col>
            <Col>
              <Form.Control
                type="text"
                name="lastName"
                placeholder="Last Name"
                defaultValue={formData.lastName}
              />
            </Col>
          </Row>
        </Form.Group>
        <Form.Group>
          <Form.Control
            type="email"
            name="email"
            placeholder="Email"
            defaultValue={formData.email}
          />
        </Form.Group>
        <Form.Group>
          <Row>
            <Col>
              <Form.Control
                type="password"
                name="password"
                placeholder="Password"
                defaultValue={formData.password}
              />
            </Col>
            <Col>
              <Form.Control
                type="password"
                name="repeatPassword"
                placeholder="Repeat password"
                defaultValue={formData.repeatPassword}
              />
            </Col>
          </Row>
        </Form.Group>
        <Button variant="primary" type="submit">
          {!signUpLoading ? (
            "Sign Up"
          ) : (
            <Spinner animation="border" size="sm" />
          )}
        </Button>
      </Form>
    </div>
  );
}

function initialForm() {
  return {
    name: "",
    lastName: "",
    email: "",
    password: "",
    repeatPassword: "",
  };
}
