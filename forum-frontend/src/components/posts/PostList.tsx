import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { postApi } from '../../services/postApi';
import type { PostResponseDTO, ErrorResponse } from '../../types';
import PostForm from './PostForm';
import PostEditForm from './PostEditForm';
import type { AxiosError } from 'axios';
import './Post.css';

interface PostListProps {
  topicId: number;
  topicTitle: string;
}

const PostList: React.FC<PostListProps> = ({ topicId, topicTitle }) => {
  const navigate = useNavigate();
  const [posts, setPosts] = useState<PostResponseDTO[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [editingPost, setEditingPost] = useState<PostResponseDTO | null>(null);
  const [userId, setUserId] = useState<number | null>(null);

  useEffect(() => {
    const fetchPosts = async () => {
      setLoading(true);
      setError('');
      try {
        const response = await postApi.getByTopic(topicId);
        setPosts(response.data || []); // always default to empty array

        const token = localStorage.getItem('token');
        if (token) {
          const payload = JSON.parse(atob(token.split('.')[1]));
          setUserId(payload.user_id ?? payload.sub ?? null);
        }
      } catch (err) {
        const axiosErr = err as AxiosError<ErrorResponse>;
        setError(axiosErr.response?.data?.error || 'Failed to load posts');
      } finally {
        setLoading(false);
      }
    };

    if (topicId) fetchPosts();
  }, [topicId]);

  const handleNewPost = (post: PostResponseDTO) => {
    setPosts(prev => [post, ...prev]);
  };

  const handleEdit = (post: PostResponseDTO) => setEditingPost(post);

  const handleUpdate = (updatedPost: PostResponseDTO) => {
    setPosts(prev => prev.map(p => (p.id === updatedPost.id ? updatedPost : p)));
    setEditingPost(null);
  };

  const handleDelete = async (id: number) => {
    if (!window.confirm('Are you sure you want to delete this post?')) return;

    try {
      await postApi.delete(id);
      setPosts(prev => prev.filter(p => p.id !== id));
    } catch (err) {
      const axiosErr = err as AxiosError<ErrorResponse>;
      alert(axiosErr.response?.data?.error || 'Failed to delete post');
    }
  };

  if (loading) return <p>Loading posts...</p>;
  if (error) return <p>{error}</p>;

  return (
    <div className="post-section">
      <h2>Posts in {topicTitle}</h2>
      {posts.length === 0 && (
        <p style={{ fontStyle: 'italic', marginBottom: '16px' }}>
          No posts yet. Be the first to create one!
        </p>
      )}

 <ul className="post-list">
  {posts.map(post => (
    <li
      key={post.id}
      className="post-card"
      onClick={() => navigate(`/posts/${post.id}`)}
    >
      {editingPost?.id === post.id ? (
        <PostEditForm
          post={editingPost}
          onCancel={() => setEditingPost(null)}
          onUpdate={handleUpdate}
        />
      ) : (
        <>
          <h3
            style={{
              margin: 0,
              cursor: 'pointer',
              color: '#0066cc',     
              textDecoration: 'underline' 
            }}
          >
            {post.title}
          </h3>
          <p>{post.content}</p>
          <small>
            By {post.username} | Created at {post.created_at}
          </small>
          {userId === post.user_id && (
            <div
              className="post-actions"
              onClick={e => e.stopPropagation()}
            >
              <button onClick={() => handleEdit(post)}>Edit</button>
              <button onClick={() => handleDelete(post.id)}>Delete</button>
            </div>
          )}
        </>
      )}
    </li>
  ))}
</ul>
      <div style={{ marginTop: '40px' }}>
        <h3>Create New Post</h3>
        <PostForm topicId={topicId} onPostCreated={handleNewPost} />
      </div>
    </div>
  );
};

export default PostList;