import React, { useState } from 'react';
import { commentApi } from '../../services/commentApi';
import type {
  CreateCommentDTO,
  CommentResponseDTO,
  ErrorResponse
} from '../../types';
import type { AxiosError } from 'axios';
import './Comment.css';

interface CommentFormProps {
  postId: number;
  parentId?: number | null;
  onCommentCreated: (comment: CommentResponseDTO) => void;
}

const CommentForm: React.FC<CommentFormProps> = ({
  postId,
  parentId = null,
  onCommentCreated
}) => {
  const [formData, setFormData] = useState<CreateCommentDTO>({
    content: '',
    post_id: postId,
    parent_id: parentId || undefined
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleChange = (
    e: React.ChangeEvent<HTMLTextAreaElement>
  ) => {
    setFormData({ ...formData, content: e.target.value });
  };

const handleSubmit = async (e: React.FormEvent) => {
  e.preventDefault();
  setLoading(true);
  setError('');

  if (!formData.content || formData.content.trim().length === 0) {
    setError('Content cannot be empty');
    setLoading(false);
    return;
  }

  try {
    let response;
    if (parentId) {
      response = await commentApi.createReply(parentId, { content: formData.content });
    } else {
      response = await commentApi.create(postId, {
        content: formData.content,
        parent_id: undefined,
        post_id: postId
      });
    }

    onCommentCreated(response.data);
    setFormData({ content: '', post_id: postId, parent_id: parentId || undefined });
  } catch (err) {
    const axiosErr = err as AxiosError<ErrorResponse>;
    setError(axiosErr.response?.data?.error || 'Failed to create comment');
  } finally {
    setLoading(false);
  }
};

  return (
    <form onSubmit={handleSubmit} style={{ marginTop: '12px' }}>
      {error && <div style={{ color: 'red' }}>{error}</div>}
      <textarea
        value={formData.content}
        onChange={handleChange}
        placeholder={parentId ? 'Write a reply...' : 'Write a comment...'}
        rows={3}
        style={{ width: '100%', marginBottom: '6px' }}
      />
      <button type="submit" disabled={loading}>
        {loading ? 'Posting...' : parentId ? 'Reply' : 'Comment'}
      </button>
    </form>
  );
};

export default CommentForm;