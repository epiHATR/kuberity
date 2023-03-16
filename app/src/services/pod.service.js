import axios from 'axios';
import authHeader from "../helpers/auth-header";

const headers = authHeader()

class PodsService {
  async getAllPods() {
    return await axios.get("/pods", { headers: headers }).then((response) => {
      return response.data.result;
    });
  }
}

export default new PodsService();
