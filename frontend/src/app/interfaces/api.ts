export interface ApiResponse<T> {
  status: string;
  data: T;
}

export interface UserSession {
  id: number;
  name: string;
  role: 'admin' | 'group_lead' | 'group_worker';
}

export type LoginResponse = ApiResponse<UserSession>;
