export interface SuccessResponse<K extends string, T> {
  success: true;
  failure: false;
  data: { [key in K]?: T | null };
  message: string;
}

export interface FailureResponse {
  success: false;
  failure: true;
  error: string;
}
