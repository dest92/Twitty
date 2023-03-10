export function isEmailValid(email) {
  // eslint-disable-next-line no-useless-escape
  const emailValid = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  return emailValid.test(String(email).toLowerCase());
}
