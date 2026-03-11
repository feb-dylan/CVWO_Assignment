import React, { useState } from 'react';
import { postApi } from '../../services/postApi';
import type { PostResponseDTO, UpdatePostDTO, ErrorResponse } from '../../types';
import type { AxiosError } from 'axios';
import './Post.css'

interface PostEditFormProps {
  post: PostResponseDTO;
  onUpdate: (updatedPost: PostResponseDTO) => void;
  onCancel: () => void;
}

const PostEditForm: React.FC<PostEditFormProps> = ({ post, onUpdate, onCancel }) => {
  const [formData, setFormData] = useState<UpdatePostDTO>({
    title: post.title,
    content: post.content
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    if (formData.title && formData.title.length < 3) {
      setError('Title must be at least 3 characters');
      setLoading(false);
      return;
    }

    if (formData.content && formData.content.length < 10) {
      setError('Content must be at least 10 characters');
      setLoading(false);
      return;
    }

    if (formData.title === post.title && formData.content === post.content) {
      onCancel();
      return;
    }

    try {
      const response = await postApi.update(post.id, formData);
      onUpdate(response.data);
    } catch (err) {
      const axiosErr = err as AxiosError<ErrorResponse>;
      setError(axiosErr.response?.data?.error || 'Failed to update post');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="post-edit-form">
      {error && <div className="form-error">{error}</div>}
      
      <div>
        <label>Title:</label>
        <input
          type="text"
          name="title"
          value={formData.title || ''}
          onChange={handleChange}
          required
          minLength={3}
          maxLength={255}
        />
      </div>

      <div>
        <label>Content:</label>
        <textarea
          name="content"
          value={formData.content || ''}
          onChange={handleChange}
          required
          minLength={10}
          rows={6}
        />
      </div>

      <div className="form-actions">
        <button type="submit" disabled={loading}>
          {loading ? 'Updating...' : 'Update Post'}
        </button>
        <button type="button" onClick={onCancel} disabled={loading}>
          Cancel
        </button>
      </div>
    </form>
  );
};

export default PostEditForm;