interface ResponseWithStatus {
  status: 'ok' | 'error';
}

export interface ApiResponse<T> extends ResponseWithStatus {
  data: T;
}

export interface ApiErrorResponse extends ResponseWithStatus {
  error_detail: string;
}

export type UserRole = 'admin' | 'group_lead' | 'group_worker';

export interface UserSession {
  id: number;
  name: string;
  role: UserRole;
}

export type SessionResponse = ApiResponse<UserSession>;
