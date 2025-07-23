import { apiV1, type FailureResponse, type SuccessResponse } from '$lib';
import { apiRequest } from './utils/clients';

export async function fetchAppointments() {
  return await apiRequest<SuccessResponse<'appointments', { State: string }[]>, FailureResponse>({
    instance: apiV1,
    path: 'appointments',
    success: response => {
      console.log(response.data);
      return {
        success: true,
        failure: false,
        message: response.data.data.message,
        data: response.data.data,
      };
    },
    failure: error => {
      console.error('Error fetching appointments:', error);
      return {
        success: false,
        failure: true,
        error: error.response.error,
      };
    },
  });
}
