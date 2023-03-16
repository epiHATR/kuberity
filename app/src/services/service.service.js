import axios from 'axios';
import authHeader from "../helpers/auth-header";

const headers = authHeader()

class ServicesService {
  async getAllServices() {
    return await axios.get("/services", { headers: headers }).then((response) => {
      return response.data.result;
    });
  }
}

export default new ServicesService();
