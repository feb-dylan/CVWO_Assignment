import api from './api';
import type {
  CommentResponseDTO,
  CreateCommentDTO,
  UpdateCommentDTO,
} from '../types';

export const commentApi = {
  getByPost: (postId: number) =>
    api.get<CommentResponseDTO[]>(`/posts/${postId}/comments`),

  getById: (id: number) =>
    api.get<CommentResponseDTO>(`/comments/${id}`),

  create: (postId: number, data: CreateCommentDTO) =>
    api.post<CommentResponseDTO>(`/posts/${postId}/comments`, data),

  update: (id: number, data: UpdateCommentDTO) =>
    api.put<CommentResponseDTO>(`/comments/${id}`, data),

  delete: (id: number) =>
    api.delete(`/comments/${id}`),

  createReply: (commentId: number, data: { content: string }) =>
    api.post<CommentResponseDTO>(`/comments/${commentId}/replies`, data),

  getReplies: (commentId: number) =>
    api.get<CommentResponseDTO[]>(`/comments/${commentId}/replies`),
};