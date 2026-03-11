import React, { useState } from 'react';
import type { CommentResponseDTO } from '../../types';
import CommentForm from './CommentForm';
import CommentEditForm from './CommentEditForm';
import { commentApi } from '../../services/commentApi';
import './Comment.css';

interface CommentItemProps {
  comment: CommentResponseDTO;
  currentUserId?: number;
  onCommentUpdated: (comment: CommentResponseDTO) => void;
  onCommentDeleted: (id: number) => void;
}

const CommentItem: React.FC<CommentItemProps> = ({
  comment,
  currentUserId,
  onCommentUpdated,
  onCommentDeleted
}) => {
  const [editing, setEditing] = useState(false);
  const [showReplyForm, setShowReplyForm] = useState(false);
  const [replies, setReplies] = useState<CommentResponseDTO[]>(comment.replies || []);

  const handleUpdate = async (content: string) => {
    try {
      const response = await commentApi.update(comment.id, { content });
      onCommentUpdated(response.data);
      setEditing(false);
    } catch (err) {
      console.error('Failed to update comment', err);
    }
  };

  const handleDelete = async () => {
    if (!confirm('Delete this comment?')) return;
    try {
      await commentApi.delete(comment.id);
      onCommentDeleted(comment.id);
    } catch (err) {
      console.error('Failed to delete comment', err);
    }
  };

  const handleReplyCreated = (reply: CommentResponseDTO) => {
    setReplies(prev => [...prev, reply]);
  };

  return (
    <div style={{ marginLeft: comment.parent_id ? 20 : 0, marginBottom: '12px' }}>
      {editing ? (
        <CommentEditForm
          initialContent={comment.content}
          onSave={handleUpdate}
          onCancel={() => setEditing(false)}
        />
      ) : (
        <>
          <div>
            <strong>{comment.username}</strong> • <small>{comment.created_at}</small>
          </div>
          <div>{comment.content}</div>
          <div style={{ marginTop: '4px' }}>
            <button onClick={() => setShowReplyForm(!showReplyForm)}>Reply</button>
            {currentUserId === comment.user_id && (
              <>
                <button onClick={() => setEditing(true)} style={{ marginLeft: '6px' }}>Edit</button>
                <button onClick={handleDelete} style={{ marginLeft: '6px' }}>Delete</button>
              </>
            )}
          </div>
        </>
      )}

      {showReplyForm && (
        <CommentForm
          postId={comment.post_id}
          parentId={comment.id}
          onCommentCreated={handleReplyCreated}
        />
      )}

      {replies.map(reply => (
        <CommentItem
          key={reply.id}
          comment={reply}
          currentUserId={currentUserId}
          onCommentUpdated={updated =>
            setReplies(prev => prev.map(r => r.id === updated.id ? updated : r))
          }
          onCommentDeleted={id =>
            setReplies(prev => prev.filter(r => r.id !== id))
          }
        />
      ))}
    </div>
  );
};

export default CommentItem;