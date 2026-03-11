import React, { useState } from 'react';
import { topicApi } from '../../services/topicApi';
import type { CreateTopicDTO, TopicResponseDTO, ErrorResponse } from '../../types';
import type { AxiosError } from 'axios';
import './Topic.css';

interface TopicFormProps {
  onTopicCreated: (topic: TopicResponseDTO) => void;
}

const TopicForm: React.FC<TopicFormProps> = ({ onTopicCreated }) => {
  const [formData, setFormData] = useState<CreateTopicDTO>({
    title: '',
    description: ''
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

    try {
      const response = await topicApi.create(formData);
      onTopicCreated(response.data);
      setFormData({ title: '', description: '' }); // reset form
    } catch (err) {
      const axiosErr = err as AxiosError<ErrorResponse>;
      setError(axiosErr.response?.data?.error || 'Failed to create topic');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="topic-form">
      <h3>Create New Topic</h3>
      {error && <div className="form-error">{error}</div>}
      <form onSubmit={handleSubmit}>
        <div>
          <label>Title</label>
          <input
            type="text"
            name="title"
            value={formData.title}
            onChange={handleChange}
            placeholder="Topic title"
            required
          />
        </div>
        <div>
          <label>Description</label>
          <textarea
            name="description"
            value={formData.description}
            onChange={handleChange}
            placeholder="Optional description"
          />
        </div>
        <button type="submit" disabled={loading}>
          {loading ? 'Creating...' : 'Create Topic'}
        </button>
      </form>
    </div>
  );
};

export default TopicForm;