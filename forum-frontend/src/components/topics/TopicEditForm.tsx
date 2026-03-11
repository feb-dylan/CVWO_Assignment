import React, { useState } from 'react';
import { topicApi } from '../../services/topicApi';
import type { TopicResponseDTO, UpdateTopicDTO, ErrorResponse } from '../../types';
import type { AxiosError } from 'axios';
import './Topic.css';

interface TopicEditFormProps {
  topic: TopicResponseDTO;
  onUpdate: (topic: TopicResponseDTO) => void;
  onCancel: () => void;
}

const TopicEditForm: React.FC<TopicEditFormProps> = ({ topic, onUpdate, onCancel }) => {
  const [formData, setFormData] = useState<UpdateTopicDTO>({
    title: topic.title,
    description: topic.description
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

    try {
      const response = await topicApi.update(topic.id, formData);
      onUpdate(response.data);
    } catch (err) {
      const axiosErr = err as AxiosError<ErrorResponse>;
      setError(axiosErr.response?.data?.error || 'Failed to update topic');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="topic-edit-form">
      {error && <div className="form-error">{error}</div>}
      <div>
        <input
          type="text"
          name="title"
          value={formData.title || ''}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <textarea
          name="description"
          value={formData.description || ''}
          onChange={handleChange}
        />
      </div>
      <button type="submit" disabled={loading}>
        {loading ? 'Updating...' : 'Update'}
      </button>
      <button type="button" onClick={onCancel} disabled={loading}>
        Cancel
      </button>
    </form>
  );
};

export default TopicEditForm;