import api from './api';
import type {
  PostResponseDTO,
  CreatePostDTO,
  UpdatePostDTO,
} from '../types';

export const postApi = {
  getAll: () =>
    api.get<PostResponseDTO[]>('/posts'),

  getByTopic: (topicId: number) =>
    api.get<PostResponseDTO[]>(`/topics/${topicId}/posts`),

  getById: (id: number) =>
    api.get<PostResponseDTO>(`/posts/${id}`),

  create: (data: CreatePostDTO) =>
    api.post<PostResponseDTO>('/posts/', data),

  update: (id: number, data: UpdatePostDTO) =>
    api.put<PostResponseDTO>(`/posts/${id}`, data),

  delete: (id: number) =>
    api.delete(`/posts/${id}`),
};