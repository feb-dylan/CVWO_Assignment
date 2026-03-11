import React from 'react';
import TopicList from '../components/topics/TopicList';

const TopicsPage: React.FC = () => {
  return (
    <div style={{ maxWidth: '900px', margin: '0 auto', padding: '20px' }}>
      <TopicList />
    </div>
  );
};

export default TopicsPage;