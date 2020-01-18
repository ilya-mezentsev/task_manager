import {ApiResponse} from './api';

export interface Group {
  name: string;
}

export type GroupsListResponse = ApiResponse<Group[]>;
