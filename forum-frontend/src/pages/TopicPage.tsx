import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { topicApi } from '../services/topicApi';
import PostList from '../components/posts/PostList';
import type { TopicResponseDTO, ErrorResponse } from '../types';
import type { AxiosError } from 'axios';

const TopicPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const [topic, setTopic] = useState<TopicResponseDTO | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!id) return;

    const fetchTopic = async () => {
      setLoading(true);
      setError('');
      try {
        const response = await topicApi.getById(Number(id));
        setTopic(response.data);
      } catch (err) {
        const axiosErr = err as AxiosError<ErrorResponse>;
        setError(axiosErr.response?.data?.error || 'Failed to load topic');
      } finally {
        setLoading(false);
      }
    };

    fetchTopic();
  }, [id]);

  if (!id) return <p>Topic not found</p>;
  if (loading) return <p>Loading topic...</p>;
  if (error) return <p>{error}</p>;
  if (!topic) return <p>Topic not found</p>;

  return (
    <div className="topic-page-container">
      <button className="back-button" onClick={() => navigate('/topics')}>
        ← Back to Topics
      </button>

      <div className="topic-card">
        <h2>{topic.title}</h2>
        {topic.description && <p>{topic.description}</p>}
        <small>
          By {topic.username} | Created at {topic.created_at}
        </small>
      </div>
      <PostList topicId={topic.id} topicTitle={topic.title} />
    </div>
  );
};

export default TopicPage;