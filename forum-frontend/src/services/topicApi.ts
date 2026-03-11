import api from './api';
import type { TopicResponseDTO, CreateTopicDTO, UpdateTopicDTO } from '../types';

export const topicApi = {
  getAll: () => api.get<TopicResponseDTO[]>('/topics'),

  getById: (id: number) => api.get<TopicResponseDTO>(`/topics/${id}`),

  create: (data: CreateTopicDTO) => api.post<TopicResponseDTO>('/topics/', data),

  update: (id: number, data: UpdateTopicDTO) => api.put<TopicResponseDTO>(`/topics/${id}`, data),

  delete: (id: number) => api.delete<{ message: string }>(`/topics/${id}`),
};