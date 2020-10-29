import { setAuthorization } from "./middleware";

export function login(credentials) {
  return new Promise((resolve, reject) => {
    window.axios
      .post("/auth/login", credentials)
      .then(response => {
        setAuthorization(response.data.access_token);
        resolve(response.data);
      })
      .catch(error => {
        reject(error);
      });
  });
}

export function getLocalUser() {
  const userStr = localStorage.getItem("user");
  if (!userStr) {
    return null;
  }
  return JSON.parse(userStr);
}
