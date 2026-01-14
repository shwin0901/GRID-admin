import { api } from '@/lib/api'
import type {
  ApiResponse,
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  User,
} from '@/types/auth'

export async function login(data: LoginRequest): Promise<LoginResponse> {
  const response = await api.post<ApiResponse<LoginResponse>>('/auth/login', data)
  return response.data.data!
}

export async function register(data: RegisterRequest): Promise<User> {
  const response = await api.post<ApiResponse<User>>('/auth/register', data)
  return response.data.data!
}

export async function getProfile(): Promise<User> {
  const response = await api.get<ApiResponse<User>>('/user/profile')
  return response.data.data!
}
