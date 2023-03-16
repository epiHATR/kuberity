import axios from 'axios';
import authHeader from "../helpers/auth-header";

const headers = authHeader()

class DeploymentsService {
  async getAllDeployments() {
    return await axios.get("/deployments", { headers: headers }).then((response) => {
      return response.data.result;
    });
  }
}

export default new DeploymentsService();
