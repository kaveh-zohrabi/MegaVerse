export interface APIResponse<T = unknown> {
  success: boolean;
  data?: T;
  error?: APIError;
  meta?: PaginationMeta;
}

export interface APIError {
  code: string;
  message: string;
}

export interface PaginationMeta {
  page: number;
  limit: number;
  total: number;
}

export interface PaginationQuery {
  page?: number;
  limit?: number;
}

export function successResponse<T>(data: T, meta?: PaginationMeta): APIResponse<T> {
  return { success: true, data, meta };
}

export function errorResponse(code: string, message: string): APIResponse<never> {
  return { success: false, error: { code, message } };
}
