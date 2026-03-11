import React, { useState } from 'react';
import { postApi } from '../../services/postApi';
import type { CreatePostDTO, PostResponseDTO, ErrorResponse } from '../../types';
import type { AxiosError } from 'axios';
import './Post.css'

interface PostFormProps {
  topicId: number;
  onPostCreated: (post: PostResponseDTO) => void;
}

const PostForm: React.FC<PostFormProps> = ({ topicId, onPostCreated }) => {
  const [formData, setFormData] = useState<CreatePostDTO>({
    title: '',
    content: '',
    topic_id: topicId
  });
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    if (formData.title.length < 3) {
      setError('Title must be at least 3 characters');
      setLoading(false);
      return;
    }

    if (formData.content.length < 10) {
      setError('Content must be at least 10 characters');
      setLoading(false);
      return;
    }

    try {
      const response = await postApi.create(formData);
      onPostCreated(response.data);
      setFormData({ title: '', content: '', topic_id: topicId }); // reset form
    } catch (err) {
      const axiosErr = err as AxiosError<ErrorResponse>;
      setError(axiosErr.response?.data?.error || 'Failed to create post');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="post-form">
      <h3>Create New Post</h3>
      {error && <div className="form-error">{error}</div>}
      <form onSubmit={handleSubmit}>
        <div>
          <label>Title</label>
          <input
            type="text"
            name="title"
            value={formData.title}
            onChange={handleChange}
            placeholder="Post title"
            required
            minLength={3}
            maxLength={255}
          />
        </div>
        <div>
          <label>Content</label>
          <textarea
            name="content"
            value={formData.content}
            onChange={handleChange}
            placeholder="Post content"
            required
            minLength={10}
            rows={6}
          />
        </div>
        <button type="submit" disabled={loading}>
          {loading ? 'Creating...' : 'Create Post'}
        </button>
      </form>
    </div>
  );
};

export default PostForm;