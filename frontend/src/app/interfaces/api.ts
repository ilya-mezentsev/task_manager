export interface ApiResponse<T> {
  status: string;
  data: T;
}

export type LoginResponse = ApiResponse<UserSession>;

export interface UserSession {
  id: number;
  name: string;
  role: string;
}
