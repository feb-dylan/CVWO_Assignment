import React, { useEffect, useState } from 'react';
import type { CommentResponseDTO } from '../../types';
import { commentApi } from '../../services/commentApi';
import CommentItem from './CommentItem';
import CommentForm from './CommentForm';

interface CommentListProps {
  postId: number;
}

const CommentList: React.FC<CommentListProps> = ({ postId }) => {
  const [comments, setComments] = useState<CommentResponseDTO[]>([]);
  const [currentUserId, setCurrentUserId] = useState<number | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string>('');

  useEffect(() => {
    const fetchComments = async () => {
      setLoading(true);
      setError('');
      // Get current user ID from token
      const token = localStorage.getItem('token');
      if (token) {
        try {
          const payload = JSON.parse(atob(token.split('.')[1]));
          setCurrentUserId((payload.user_id ?? payload.sub) || null);
        } catch {
          setCurrentUserId(null);
        }
      }

      try {
        const response = await commentApi.getByPost(postId);
        setComments(response.data || []);
      } catch {
        setComments([]);
        setError('Failed to load comments');
      } finally {
        setLoading(false);
      }
    };

    fetchComments();
  }, [postId]);

  const handleCommentCreated = (comment: CommentResponseDTO) => {
    setComments(prev => [comment, ...prev]);
  };

  const handleCommentUpdated = (updated: CommentResponseDTO) => {
    setComments(prev => prev.map(c => (c.id === updated.id ? updated : c)));
  };

  const handleCommentDeleted = (id: number) => {
    setComments(prev => prev.filter(c => c.id !== id));
  };

  return (
    <div style={{ marginTop: '16px' }}>
      {/* Always show comment form at top */}
      <CommentForm postId={postId} onCommentCreated={handleCommentCreated} />

      {loading && <p>Loading comments...</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}

      {!loading && !error && comments.length === 0 && (
        <p style={{ fontStyle: 'italic', marginTop: '12px' }}>
          No comments yet. Be the first to comment!
        </p>
      )}

      {!loading && !error && comments.map(comment => (
        <div
          key={comment.id}
          style={{
            border: '1px solid #ddd',
            borderRadius: '6px',
            padding: '12px',
            marginBottom: '8px',
            backgroundColor: '#fafafa',
            boxShadow: '0 1px 3px rgba(0,0,0,0.05)'
          }}
        >
          <CommentItem
            comment={comment}
            currentUserId={currentUserId ?? undefined}
            onCommentUpdated={handleCommentUpdated}
            onCommentDeleted={handleCommentDeleted}
          />
        </div>
      ))}
    </div>
  );
};

export default CommentList;