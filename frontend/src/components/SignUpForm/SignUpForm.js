import React from "react";
import "./SignUpForm.scss";
import { Row, Col, Form, Button, Spinner } from "react-bootstrap";

export default function SignUpForm(props) {
  const { setShowModal } = props;

  const onSubmit = (e) => {
    e.preventDefault(); //Prevents page refresh
    setShowModal(false);
  };

  return (
    <div className="sign-up-form">
      <h2>Create your account</h2>
      <Form onSubmit={onSubmit}>
        <Form.Group>
          <Row>
            <Col>
              <Form.Control type="text" placeholder="Name" />
            </Col>
            <Col>
              <Form.Control type="text" placeholder="Last Name" />
            </Col>
          </Row>
        </Form.Group>
        <Form.Group>
          <Form.Control type="email" placeholder="Email" />
        </Form.Group>
        <Form.Group>
            <Row> 
            <Col>
            <Form.Control type="password" placeholder="Password" />
            </Col>
            <Col>
            <Form.Control type="password" placeholder="Repeat password" />
            </Col>
            </Row>
        </Form.Group>
        <Button variant="primary" type="submit">
          Sign up
        </Button>
      </Form>
    </div>
  );
}
