import axios from 'axios';

class AuthService {
  async login(user) {
    let payload = {
      username: user.username,
      password: user.password,
    };
    const response = await axios.post('/account/verify', payload);
    return response.data;
  }
}

export default new AuthService();
