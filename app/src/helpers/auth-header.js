export default function authHeader() {
  let user = JSON.parse(localStorage.getItem("user"));
  if (user && user.accesstoken) {
    return {
      Authorization: "Bearer " + user.accesstoken,
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*"
    };
  } else {
    return {};
  }
}
