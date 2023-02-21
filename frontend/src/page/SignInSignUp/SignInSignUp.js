import React from "react";
import { Container, Row, Col, Button } from "react-bootstrap";

export default function SignInSignUp() {
  return (
    <Container className="signin-signup" fluid>
      <Row>
        <LeftComponent />
        <RightComponent />
      </Row>
    </Container>
  );
}

function LeftComponent() {
  return (
    <Col className="signin-signup-left">
      <h2>Left Component</h2>
    </Col>
  );
}

function RightComponent() {
    return (
        <Col className="signin-signup-right">
        <h2>Right Component</h2>
        </Col>
    );
}

