import api from './api';
import type { LoginDTO, RegisterDTO, UserResponseDTO } from '../types';

export interface LoginResponse {
  message: string;
  token: string;
  user: UserResponseDTO;
}

export interface RegisterResponse {
  message: string;
  user: UserResponseDTO;
}

export const authApi = {
  register: (data: RegisterDTO) =>
    api.post<RegisterResponse>('/auth/register', data),

  login: (data: LoginDTO) =>
    api.post<LoginResponse>('/auth/login', data),

  logout: () => {
    localStorage.removeItem('token');
  },
};