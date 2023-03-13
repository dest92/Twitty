import React, { useState } from "react";
import { Container, Row, Col, Button } from "react-bootstrap";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faSearch,
  faUsers,
  faComment,
} from "@fortawesome/free-solid-svg-icons";
import BasicModal from "../../components/Modal/BasicModal";
import SignUpForm from "../../components/SignUpForm";
import SignInForm from "../../components/SignInForm";
import WhiteLogoTwitty from "../../assets/png/logo-white.png";
import LogoTwitty from "../../assets/png/logo.png";
import "./SignInSignUp.scss";

export default function SignInSignUp() {
  const [showModal, setShowModal] = useState(false);
  const [contentModal, setModalContent] = useState(null);

  // Define a function to open the modal, which sets the modal content and shows the modal
  const openModal = (content) => {
    setModalContent(content);
    setShowModal(true);
  };

  return (
    <>
      <Container className="signin-signup" fluid>
        <Row>
          <LeftComponent />
          <RightComponent
            openModal={openModal}
            setShowModal={setShowModal}
            // setRefreshCheckLogin={setRefreshCheckLogin}
          />
        </Row>
      </Container>

      <BasicModal show={showModal} setShow={setShowModal}>
        {contentModal}
        {/* Display the content of the modal depending of the value of content Modal */}
      </BasicModal>
    </>
  );
}

function LeftComponent() {
  return (
    <Col className="signin-signup__left" xs={6}>
      <img src={LogoTwitty} alt="Twitty!" />
      <div>
        <h2>
          <FontAwesomeIcon icon={faSearch} />
          Follow any what do you want
        </h2>
        <h2>
          <FontAwesomeIcon icon={faUsers} />
          Hear what people are talking about
        </h2>
        <h2>
          <FontAwesomeIcon icon={faComment} />
          Join the conversation
        </h2>
      </div>
    </Col>
  );
}

function RightComponent(props) {
  const { openModal, setShowModal } = props;

  return (
    <Col className="signin-signup__right" xs={6}>
      <div>
        <h2>See what's happening in the world right now</h2>
        <h3>Join Twitty today.</h3>

        <Button
          onClick={() => openModal(<SignUpForm setShowModal={setShowModal} />)}
          variant="primary"
        >
          Sign up
        </Button>

        <Button
          onClick={() => openModal(<SignInForm />)}
          variant="outline-primary"
        >
          Sign in
        </Button>
      </div>
    </Col>
  );
}
