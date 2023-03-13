import axios from "axios";
import { API_HOST } from "../utils/constant";

export async function signUpApi(user) {
  if (!navigator.onLine) {
    throw new Error("No internet connection");
  }

  const endpoint = "register";
  const contentType = "application/json";

  const userTemp = {
    ...user,
    email: user.email.toLowerCase(),
    birthDate: new Date(),
  };

  delete userTemp.repeatPassword;

  const url = `${API_HOST}/${endpoint}`;
  const data = JSON.stringify(userTemp);

  const config = {
    headers: {
      "Content-Type": contentType,
    },
  };

  try {
    const response = await axios.post(url, data, config);
    if (response.status >= 200 && response.status < 300) {
      return response.data;
    } else {
      return { code: 404, message: "Email no available" };
    }
  } catch (err) {
    if (err.code === "ERR_CONNECTION_REFUSED") {
      return { code: 500, message: "Error: ERR_CONNECTION_REFUSED" };
    } else {
      throw new Error("Error from server, try again later!");
    }
  }
}
