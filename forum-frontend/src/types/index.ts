export interface UserResponseDTO {
  id: number;
  username: string;
}

export interface TopicResponseDTO {
  id: number;
  title: string;
  description: string;
  user_id: number;
  username: string;
  created_at: string;
  updated_at: string;
}

export interface PostResponseDTO {
  id: number;
  title: string;
  content: string;
  topic_id: number;
  user_id: number;
  username: string;
  created_at: string;
  updated_at: string;
}

export interface CommentResponseDTO {
  id: number;
  content: string;
  post_id: number;
  user_id: number;
  username: string;
  parent_id: number | null;
  created_at: string;
  updated_at: string;
  replies?: CommentResponseDTO[];
}

export interface CreateTopicDTO {
  title: string;
  description: string;
}

export interface UpdateTopicDTO {
  title?: string;
  description?: string;
}

export interface CreatePostDTO {
  title: string;
  content: string;
  topic_id: number;
}

export interface UpdatePostDTO {
  title?: string;
  content?: string;
}

export interface CreateCommentDTO {
  content: string;
  parent_id?: number | null;
  post_id: number;
}

export interface UpdateCommentDTO {
  content?: string;
}

export interface RegisterDTO {
  username: string;
  password: string;
}

export interface LoginDTO {
  username: string;
  password: string;
}
export interface LoginResponse {
  message: string;
  token: string;
  user: UserResponseDTO;
}

export interface RegisterResponse {
  message: string;
  user: UserResponseDTO;
}
export interface ErrorResponse {
  error: string;
}