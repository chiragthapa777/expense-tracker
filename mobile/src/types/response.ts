export interface ApiResponse<T> {
  success: true;
  data: T;
}

export interface ApiResponsePagination<T> {
  success: true;
  data: T[];
  metadata: {
    pagination: {
      total: number;
      limit: number;
      currentPage: number;
      totalPages: number;
    };
  };
}

export interface ApiErrorResponse {
  success: true;
  error: string;
  code?: string | "1001"; // password not set;
}
