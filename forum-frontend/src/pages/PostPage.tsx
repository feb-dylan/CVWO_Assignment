import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { postApi } from '../services/postApi';
import type { PostResponseDTO, ErrorResponse } from '../types';
import CommentList from '../components/comments/CommentList';
import type { AxiosError } from 'axios';

const PostPage: React.FC = () => {
  const { postId } = useParams<{ postId: string }>();
  const navigate = useNavigate();

  const [post, setPost] = useState<PostResponseDTO | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchPost = async () => {
      if (!postId) return;
      setLoading(true);
      setError('');
      try {
        const response = await postApi.getById(Number(postId));
        setPost(response.data);
      } catch (err) {
        const axiosErr = err as AxiosError<ErrorResponse>;
        setError(axiosErr.response?.data?.error || 'Failed to load post');
      } finally {
        setLoading(false);
      }
    };

    fetchPost();
  }, [postId]);

  if (!postId) return <p>Post not found</p>;
  if (loading) return <p>Loading post...</p>;
  if (error) return <p>{error}</p>;
  if (!post) return <p>Post not found</p>;

  return (
    <div style={{ maxWidth: '900px', margin: '0 auto', padding: '20px' }}>
      {/* Back button */}
      <button
        onClick={() => navigate(-1)}
        style={{
          marginBottom: '20px',
          padding: '8px 16px',
          border: '1px solid #ddd',
          borderRadius: '4px',
          cursor: 'pointer',
          background: '#f0f0f0',
        }}
      >
        ← Back
      </button>

      {/* Post card */}
      <div style={{
        border: '1px solid #ddd',
        borderRadius: '8px',
        padding: '16px',
        marginBottom: '20px',
        backgroundColor: '#fff',
        boxShadow: '0 2px 5px rgba(0,0,0,0.05)'
      }}>
        <h2 style={{ margin: '0 0 8px 0' }}>{post.title}</h2>
        <p style={{ marginBottom: '8px' }}>{post.content}</p>
        <small style={{ color: '#555' }}>
          By {post.username} | Created at {post.created_at}
        </small>
      </div>

      <div className="comments-section" style={{ marginTop: '20px' }}>
        <h3 style={{ marginBottom: '12px' }}>Comments</h3>
        <CommentList postId={post.id} />
      </div>
    </div>
  );
};

export default PostPage;