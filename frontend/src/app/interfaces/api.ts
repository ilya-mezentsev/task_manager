interface ResponseWithStatus {
  status: ResponseStatus;
}

export enum ResponseStatus {
  Ok = 'ok',
  Error = 'error'
}

export enum UserRole {
  Admin = 'admin',
  GroupLead = 'group_lead',
  GroupWorker = 'group_worker'
}

export interface ApiResponse<T> extends ResponseWithStatus {
  data: T;
}

export interface ApiDefaultResponse extends ResponseWithStatus {
  data: null;
}

export interface ApiErrorResponse extends ResponseWithStatus {
  error_detail: string;
}

export interface UserSession {
  id: number;
  name: string;
  role: UserRole;
  group_id: number;
}

export type SessionResponse = ApiResponse<UserSession>;
