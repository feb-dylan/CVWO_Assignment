import React, { useEffect, useState } from 'react';
import { topicApi } from '../../services/topicApi';
import type { TopicResponseDTO } from '../../types';
import TopicForm from './TopicForm';
import TopicEditForm from './TopicEditForm';
import PostList from '../posts/PostList';
import './Topic.css';

const TopicList: React.FC = () => {
  const [topics, setTopics] = useState<TopicResponseDTO[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [editingTopic, setEditingTopic] = useState<TopicResponseDTO | null>(null);
  const [userId, setUserId] = useState<number | null>(null);
  const [selectedTopic, setSelectedTopic] = useState<TopicResponseDTO | null>(null);

  useEffect(() => {
    const fetchTopics = async () => {
      try {
        const response = await topicApi.getAll();
        setTopics(response.data);

        const token = localStorage.getItem('token');
        if (token) {
          const payload = JSON.parse(atob(token.split('.')[1]));
          setUserId(payload.user_id || payload.sub);
        }
      } catch {
        setError('Failed to load topics');
      } finally {
        setLoading(false);
      }
    };

    fetchTopics();
  }, []);

  const handleNewTopic = (topic: TopicResponseDTO) => setTopics(prev => [...prev, topic]);
  const handleEdit = (topic: TopicResponseDTO) => setEditingTopic(topic);
  const handleUpdate = (updatedTopic: TopicResponseDTO) => {
    setTopics(prev => prev.map(t => (t.id === updatedTopic.id ? updatedTopic : t)));
    setEditingTopic(null);
    if (selectedTopic?.id === updatedTopic.id) setSelectedTopic(updatedTopic);
  };
  const handleDelete = async (id: number) => {
    if (!confirm('Delete this topic?')) return;
    try {
      await topicApi.delete(id);
      setTopics(prev => prev.filter(t => t.id !== id));
      if (selectedTopic?.id === id) setSelectedTopic(null);
    } catch {
      alert('Failed to delete topic');
    }
  };
  const handleTopicClick = (topic: TopicResponseDTO) => setSelectedTopic(topic);
  const handleBackToTopics = () => setSelectedTopic(null);

  if (loading) return <p>Loading topics...</p>;
  if (error) return <p>{error}</p>;

  // Show selected topic with posts
  if (selectedTopic) {
    return (
      <div style={{ maxWidth: '900px', margin: 'auto' }}>
        <button onClick={handleBackToTopics} style={{ marginBottom: '20px' }}>
          ← Back to Topics
        </button>

        <h2>{selectedTopic.title}</h2>
        {selectedTopic.description && <p>{selectedTopic.description}</p>}
        <small>By {selectedTopic.username} | {selectedTopic.created_at}</small>

        {userId === selectedTopic.user_id && (
          <div style={{ marginTop: '16px' }}>
            <button onClick={() => handleEdit(selectedTopic)}>Edit Topic</button>
            <button onClick={() => handleDelete(selectedTopic.id)} style={{ marginLeft: '10px' }}>
              Delete Topic
            </button>
          </div>
        )}

        {editingTopic && (
          <TopicEditForm
            topic={editingTopic}
            onCancel={() => setEditingTopic(null)}
            onUpdate={handleUpdate}
          />
        )}

        <PostList topicId={selectedTopic.id} topicTitle={selectedTopic.title} />
      </div>
    );
  }

  // Show all topics (list view)
  return (
    <div style={{ maxWidth: '900px', margin: 'auto' }}>
      <h2>Topics</h2>

      <div>
        {topics.map(topic => (
          <div key={topic.id} style={{ border: '1px solid #ddd', borderRadius: '8px', padding: '12px 16px', marginBottom: '8px', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            {editingTopic?.id === topic.id ? (
              <TopicEditForm
                topic={editingTopic}
                onCancel={() => setEditingTopic(null)}
                onUpdate={handleUpdate}
              />
            ) : (
              <>
                <h3 onClick={() => handleTopicClick(topic)} style={{ margin: 0, cursor: 'pointer', color: '#0066cc', flex: 1 }}>
                  {topic.title}
                </h3>

                {userId === topic.user_id && (
                  <div>
                    <button onClick={() => handleEdit(topic)}>Edit</button>
                    <button onClick={() => handleDelete(topic.id)} style={{ marginLeft: '8px' }}>
                      Delete
                    </button>
                  </div>
                )}
              </>
            )}
          </div>
        ))}
      </div>

      <div style={{ marginTop: '40px' }}>
        <h3>Create New Topic</h3>
        <TopicForm onTopicCreated={handleNewTopic} />
      </div>
    </div>
  );
};

export default TopicList;