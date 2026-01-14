export interface User {
  id: number
  username: string
  email: string
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  username: string
  password: string
}

export interface RegisterRequest {
  username: string
  password: string
  email?: string
}

export interface LoginResponse {
  token: string
  user_info: User
}

export interface ApiResponse<T = unknown> {
  code: number
  message: string
  data?: T
}
