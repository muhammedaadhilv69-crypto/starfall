components {
  id: "enemy_laser"
  component: "/main/scripts/enemy_laser.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"laserRed16\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/atlases/game.atlas\"\n"
  "}\n"
  ""
  rotation {
    z: -1.0
    w: 6.123234E-17
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy_bullets\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: -20.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"enemylasercolobj\"\n"
  "  }\n"
  "  data: 6.8\n"
  "}\n"
  ""
}
